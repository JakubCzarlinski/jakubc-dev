import { addIconSelectors } from "@iconify/tailwind";
import { skeleton } from "@skeletonlabs/tw-plugin";
import { join } from "path";
import type { Config } from "tailwindcss";
import { theme } from "./theme.ts";

export default {
  darkMode: "class",
  content: [
    "./src/**/*.{html,js,svelte,ts,templ,go,cjs,mjs}",
    "./compile/**/*.{html,js,svelte,ts,templ,go,cjs,mjs}",
    join(
      require.resolve("@skeletonlabs/skeleton"),
      "./src/**/*.{html,js,svelte,ts,templ,go,cjs,mjs}",
    ),
  ],
  theme: {
    extend: {
      transitionTimingFunction: {
        expo: "cubic-bezier(0.16, 1, 0.3, 1)",
      },
      maxWidth: {
        "3/4": "75%",
      },
    },
  },
  plugins: [
    addIconSelectors(["ic"]),
    skeleton({
      themes: {
        custom: [theme],
      },
    }),
  ],
} as Config;
