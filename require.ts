import { createRequire } from "module";

const require = createRequire(import.meta.url)("path");

function resolve(...paths: string[]) {
  return require.resolve(...paths);
}

export { resolve };
