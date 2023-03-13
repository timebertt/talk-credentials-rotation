import Reveal from 'reveal.js/dist/reveal.esm.js';

// plugins
import Markdown from 'reveal.js/plugin/markdown/markdown.esm.js';
import Highlight from 'reveal.js/plugin/highlight/highlight.esm.js';
import Notes from 'reveal.js/plugin/notes/notes.esm.js';
import QRCode from './reveal/plugin/slides-qr-code.js';

// styles
import './custom.css';
import 'reveal.js/dist/reset.css';
import 'reveal.js/dist/reveal.css';
import 'reveal.js/dist/theme/black.css';
import 'highlight.js/styles/obsidian.css';

// import all markdown files from ./content
// this loads them as webpack assets that our HTML template transforms into individual reveal sections
(ctx => {
  return ctx.keys().map(ctx);
})(require.context('./content/', false, /\.md$/));

// initialize reveal.js and plugins
Reveal.initialize({
  plugins: [Markdown, Highlight, Notes, QRCode],
  hash: true,
  history: true,
  center: false,
  controls: true,
  navigationMode: 'default',
  slideNumber: true
});
