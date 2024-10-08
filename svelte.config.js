import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";
import { removeSSR } from "./rendering/filter.ts";

// Consult https://svelte.dev/docs#compile-time-svelte-preprocess
// for more information about preprocessors
/**
 * @type {import('@sveltejs/vite-plugin-svelte').Options}
 */
export default {
  // compilerOptions: {dev: true},
  preprocess: [removeSSR(), vitePreprocess()],
};
