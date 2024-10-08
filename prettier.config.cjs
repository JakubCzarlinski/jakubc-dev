module.exports = {
  plugins: ["prettier-plugin-tailwindcss", "prettier-plugin-svelte"],
  tailwindConfig: "./tailwind.config.ts",
  endOfLine: "crlf",
  overrides: [{ files: "*.svelte", options: { parser: "svelte" } }],
};
