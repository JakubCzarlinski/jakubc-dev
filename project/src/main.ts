import "@/project/src/app.css";
import { hydrate } from "svelte";

const modules = import.meta.webpackContext("@/project/src/lib/", {
  recursive: true,
  regExp: /\.svelte$/,
});
function getComponentName(path: string) {
  return path.slice(path.lastIndexOf("/") + 1, -".svelte".length);
}

for (const path in modules.keys()) {
  const elements = document.getElementsByClassName(getComponentName(path));
  const length = elements.length;
  if (length === 0) continue;

  const module = modules(path);

  for (let i = 0; i < length; i++) {
    mountComponent(elements[i] as HTMLElement, module);
  }
}

async function mountComponent(element: HTMLElement, Component: any) {
  if (element.firstElementChild === null) return;
  const attr = element.dataset.svelte;
  if (attr === undefined) return;

  hydrate(Component, {
    target: element.firstElementChild,
    props: JSON.parse(attr),
  });
  element.replaceChildren(...element.children[0].children);
  delete element.dataset.svelte;
}

// Onpage load event
window.addEventListener("load", () => {
  const hash = window.location.hash;
  setTimeout(() => {
    if (hash !== "") {
      const element = document.getElementById(hash.slice(1));
      if (element !== null) element.scrollIntoView();
    }
  });
});
