import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

const csrRegexHTML =
  /<!--\s+CSR\s+-->([\s\S]*?)<!--\s+SSR\s+-->([\s\S]*?)<!--\s+END\s+-->/gm;
const csrRegexJS = /\/\/\s+CSR\s*([\s\S]*?)\/\/\s+SSR\s*([\s\S]*?)\/\/\s+END/gm;

/** @param {string} content */
function removeSsrInJs(content) {
  return content.replaceAll(csrRegexJS, (_, csr, _ssr) => csr);
}

/** @param {string} content */
function removeSsrInHtml(content) {
  return content.replaceAll(csrRegexHTML, (_, csr, _ssr) => csr);
}
/** @returns {import("svelte/compiler").PreprocessorGroup} */
function removeSSR() {
  return {
    markup({ content }) {
      return { code: removeSsrInJs(removeSsrInHtml(content)) };
    },
  };
}

// Consult https://svelte.dev/docs#compile-time-svelte-preprocess
// for more information about preprocessors
/** @type {import('@sveltejs/vite-plugin-svelte').Options} */
export default {
  // compilerOptions: {dev: true},
  preprocess: [removeSSR(), vitePreprocess()],
};
