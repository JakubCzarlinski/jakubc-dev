import "@/project/src/app.css";
import { hydrate } from "svelte";

const modules = import.meta.glob("@/project/src/lib/**/*.svelte");
Object.keys(modules).forEach((path) => modules[path]());

const targets: string[] = ["App", "Counter"];

for (let i = 0; i < targets.length; i++) {
  loadComponents(targets[i]);
}

async function loadComponents(className: string) {
  const elem = document.getElementsByClassName(className);
  if (elem.length === 0) return;
  const component = (await import(`/assets/${elem[0].className}.js`)).default;
  for (let i = 0; i < elem.length; i++) {
    mountComponent(elem[i], component);
  }
}

async function mountComponent(element: Element, Component: any) {
  hydrate(Component, {
    target: element.firstElementChild as Element,
    props: JSON.parse(element.getAttribute("svelte") as string),
  });
  element.replaceChildren(...element.children[0].children);
  element.removeAttribute("svelte");
}
