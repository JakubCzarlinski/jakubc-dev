<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import Project from "@/project/src/lib/Project.svelte";
  import Top from "@/project/src/lib/Top.svelte";
</script>

<CopyLink link="https://jakubc.dev/#portfolio" />

<h1 class="wavy mb-12 inline-block underline-offset-[12px]">Portfolio</h1>

<p class="lowercase">
  Portfolio coming soon. In the meantime, check out my
  <a
    class="wavy hover-text text-token"
    href="https://github.com/JakubCzarlinski/"
    target="_blank"
  >
    GitHub profile
  </a>.
</p>

<div class="h-[10vh]"></div>

<Project
  id="google"
  title="Google Internships"
  description="2022 STEP Internship in London, with a return offer for 2023 SWE Internship."
  image="https://cataas.com/cat"
/>

<Top />
<div class="h-[10vh]"></div>

<Project
  id="dissertaion"
  title="Diffusion Based Multi-Speaker TTS Synthesis"
  description="Third year dissertation project. Applying diffusion models to
generate speech from multiple speakers. Approach was similar to NaturalSpeech2,
but instead we used speaker embeddings to condition the model."
  image="https://cataas.com/cat"
  imageSide="right"
/>

<Top />
<div class="h-[10vh]"></div>

<Project
  id="framework"
  title="Go-Svelte Framework"
  description="A blazingly
fast full-stack framework that Server-Side-Renders Svelte components into Go
templates - at compile time! We get full expresivity of Svelte, and the speed of
Go. It's a match made in heaven. How fast does it build? Minimal build time is
0.6s for a full-stack app. How fast does it run? Running in a docker container,
it achieves more than 400k requests per second. In production, it uses a
stunning 6MB of RAM. On top of that, it's a single binary, and it's easy to
deploy. It's a dream come true for developers who want to build fast, secure,
and scalable web applications."
  image="https://cataas.com/cat"
/>

<Top />
<div class="h-[10vh]"></div>

<Project
  id="github"
  title="GitHub"
  link="https://www.github.com/JakubCzarlinski/"
  description="Collection of my
personal projects, open source contribution, and a showcase of my skillsets."
  image="https://cataas.com/cat"
  imageSide="right"
/>
