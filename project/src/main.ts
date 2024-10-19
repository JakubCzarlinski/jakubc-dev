import "@/project/src/app.css";
import { hydrate } from "svelte";

const modules = import.meta.glob("@/project/src/lib/**/*.svelte");
function getComponentName(path: string) {
  return path.slice(path.lastIndexOf("/") + 1, -".svelte".length);
}

for (const path in modules) {
  const elem = document.getElementsByClassName(getComponentName(path));
  const length = elem.length;
  if (length === 0) continue;

  modules[path]().then((result) => {
    for (let i = 0; i < length; i++) {
      mountComponent(elem[i], (result as { default: any }).default);
    }
  });
}

async function mountComponent(element: Element, Component: any) {
  if (element.firstElementChild === null) return;
  const attr = element.getAttribute("svelte");
  if (attr === null) return;

  hydrate(Component, {
    target: element.firstElementChild,
    props: JSON.parse(attr),
  });
  element.replaceChildren(...element.children[0].children);
  element.removeAttribute("svelte");
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
