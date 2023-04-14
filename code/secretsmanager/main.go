package main

import (
	"time"

	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/pkg/utils/secrets"
	secretsmanager "github.com/gardener/gardener/pkg/utils/secrets/manager"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/utils/clock/testing"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

// Use the GoLand run config or run this from the code directory with:
//
//	KUBECONFIG=$PWD/kind-kubeconfig.yaml go run ./secretsmanager
func main() {
	ctx := signals.SetupSignalHandler()
	log := logger.MustNewZapLogger(logger.InfoLevel, logger.FormatText)
	cl := testing.NewFakeClock(time.Date(2023, 4, 17, 11, 10, 0, 0, time.Local))

	// initialize client and prepare demo namespace
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	utilruntime.Must(err)

	namespace := "secrets-manager"
	utilruntime.Must(client.IgnoreAlreadyExists(c.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}})))

	// initialize secrets manager
	rotationTriggerTimes := map[string]time.Time{
		// "demo-ca": cl.Now(), // step 1: start CA rotation
	}

	secretsManager, err := secretsmanager.New(
		ctx,
		log.WithName("secretsmanager"),
		cl,
		c,
		namespace,
		"demo",
		secretsmanager.Config{SecretNamesToTimes: rotationTriggerTimes},
	)
	utilruntime.Must(err)

	// prepare a CA for our demo
	caSecret, err := secretsManager.Generate(ctx,
		&secrets.CertificateSecretConfig{
			CertType:   secrets.CACert,
			Name:       "demo-ca",
			CommonName: "my-selfsigned-ca",
		},
		secretsmanager.Rotate(secretsmanager.KeepOld),
		// secretsmanager.IgnoreOldSecrets(), // step 2: complete CA rotation
	)
	utilruntime.Must(err)

	log.Info("We have a CA!", "secretName", caSecret.Name)

	// BEGIN: live coding section

	serverSecret, err := secretsManager.Generate(ctx,
		&secrets.CertificateSecretConfig{
			CertType:   secrets.ServerCert,
			Name:       "demo-server",
			CommonName: "my-server",
		},
		secretsmanager.SignedByCA("demo-ca"),
	)
	utilruntime.Must(err)

	log.Info("Using server cert...", "secretName", serverSecret.Name)

	// END: live coding section

	// clean up old secrets that are no longer needed
	err = secretsManager.Cleanup(ctx)
	utilruntime.Must(err)
}
