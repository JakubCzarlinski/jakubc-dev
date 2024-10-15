<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import Project from "@/project/src/lib/Project.svelte";
  import Top from "@/project/src/lib/Top.svelte";
</script>

<CopyLink link="https://jakubc.dev/#portfolio" />

<h1 class="wavy mb-12 inline-block underline-offset-[12px]">portfolio.</h1>

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
  description=""
  image="https://cataas.com/cat"
/>

<Top />

<Project
  id="dissertaion"
  title="Diffusion Based Multi-Speaker TTS Synthesis"
  description="Third year dissertation project. Applying diffusion models to
generate speech from multiple speakers. Approach was similar to NaturalSpeech2,
but instead we used speaker embeddings to condition the model."
  image="https://cataas.com/cat"
/>

<Top />

<Project
  id="framework"
  title="Go-Svelte Framework"
  description="A blazingly
fast full-stack framework that Server-Side-Renders Svelte components into Go
templates - at compile time! We get full expresivity of Svelte, and the speed of
Go. It's a match made in heaven. How fast does it build? Minimal build time is
0.6s for a full-stack app. How fast does it run? Whilst using less RAM than notepad.exe"
  image="https://cataas.com/cat"
/>

<Top />

<Project
  id="github"
  title="GitHub"
  link="https://www.github.com/JakubCzarlinski/"
  description="Collection of my
personal projects, open source contribution, and a showcase of my skillsets."
  image="https://cataas.com/cat"
/>
