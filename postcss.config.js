import purgecss from "@fullhuman/postcss-purgecss";
import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";

function tailwindExtractor(content) {
  // Capture as liberally as possible, including things like `h-(screen-1.5)`
  const broadMatches = content.match(/[^<>"'`\s]*[^<>"'`\s:]/g) || [];
  const broadMatchesWithoutTrailingSlash = broadMatches.map((match) =>
    match.replace(/\\$/, ""),
  );

  // Capture classes within other delimiters like .block(class="w-1/2") in Pug
  const innerMatches =
    content.match(/[^<>"'`\s.(){}[\]#=%]*[^<>"'`\s.(){}[\]#=%:]/g) || [];

  return broadMatches
    .concat(broadMatchesWithoutTrailingSlash)
    .concat(innerMatches);
}

/** @type {import('postcss-load-config').Config} */
const cfg = {
  plugins: [
    tailwindcss({
      config: "./tailwind.config.ts",
    }),
    autoprefixer(),
    purgecss({
      content: ["./project/**/*.{html,js,svelte,ts,templ,go,cjs,mjs,css}"],
      extractors: [
        {
          extractor: tailwindExtractor,
          extensions: [
            "html",
            "js",
            "svelte",
            "ts",
            "templ",
            "go",
            "cjs",
            "mjs"
          ],
        },
      ],
    }),
  ],
};

export default cfg;
