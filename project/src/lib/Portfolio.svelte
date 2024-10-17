<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import Project from "@/project/src/lib/Project.svelte";
</script>

<CopyLink link="https://jakubc.dev/#portfolio" />

<h1 class="wavy mb-12 inline-block underline-offset-[12px]">Portfolio</h1>

<p class="mb-4">
  Portfolio here is work in progress. For more projects and contributions, check
  <a
    class="wavy hover-text text-token"
    href="https://github.com/JakubCzarlinski/"
    target="_blank"
  >
    GitHub profile
  </a>.
</p>

<hr class="!border-t-2" />

<Project
  id="framework"
  title="Go-Svelte Framework"
  description="A blazingly
fast full-stack framework that Server-Side Renders Svelte components into Go
templates — at compile time! It combines the full expressivity of Svelte with the
speed of Go, creating an ideal synergy. This single binary is designed to
minimize memory footprint while preserving the rapid development experience of
the JavaScript ecosystem. It's a dream come true for developers building fast,
scalable, and visually appealing web applications."
  image="https://cataas.com/cat"
  initialOpen
>
  {#snippet chips()}
    <div class="flex flex-wrap">
      <img
        src="https://img.shields.io/badge/Bun-%23000000.svg?style=for-the-badge&logo=bun&logoColor=white"
        alt="Bun"
      />
      <img
        src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white"
        alt="PostgreSQL"
      />
      <img
        src="https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white"
        alt="Vite"
      />
      <img
        src="https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white"
        alt="TypeScript"
      />
      <img
        src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"
        alt="Go"
      />
      <img
        src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white"
        alt="Docker"
      />
      <img
        src="https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white"
        alt="Tailwind CSS"
      />
      <img
        src="https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white"
        alt="Cloudflare"
      />
      <img
        src="https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white"
        alt="Svelte"
      />
      <img
        src="https://img.shields.io/badge/RollupJS-ef3335?style=for-the-badge&logo=rollup.js&logoColor=white"
        alt="RollupJS"
      />
    </div>
  {/snippet}
  {#snippet bottonSnippet()}
    <div class="text-justify leading-loose space-y-4">
      <h3 class="italic font-bold">
        Marketing fluff aside, what do the numbers say?
      </h3>
      <p class="indent-4 pl:8 md:pl-16">
        Minimal build time for a full-stack app is just
        <strong>0.6 seconds.</strong>

        Running this website inside an Alpine Docker container, the frame work
        handles over <strong>400K requests per second.</strong>

        For comparison "<code>hello world</code>" example for the Bun server can
        handle about 100k requests per second. That is 3.26GB versus 11.55MB of
        data transferred.
      </p>
      <h3 class="italic font-bold">So what about memory footprint?</h3>
      <p class="indent-4 pl:8 md:pl-16">
        Bun uses 42.8MB of memory on standby and peaks at 140MB during
        benchmarking. 42.8MB for a "<code>hello world</code>" doing nothing?!?
        Go-Svelte uses <strong>just 6.0MB and 40.1MB respectively.</strong>
      </p>
    </div>
  {/snippet}
</Project>

<Project
  id="google"
  title="Google Internships"
  description="2022 STEP Internship in London, with a return offer for 2023 SWE Internship."
  image="https://cataas.com/cat"
/>

<Project
  id="dissertaion"
  title="Diffusion Multi-Speaker TTS Synthesis"
  description="Third-year dissertation project focused on applying diffusion
models to generate speech from multiple speakers. Inspired by the NaturalSpeech2
framework, but innovated by conditioning the model with speaker embeddings for
more accurate and diverse speech generation"
  image="https://cataas.com/cat"
  imageSide="right"
/>

<Project
  id="github"
  title="GitHub"
  link="https://www.github.com/JakubCzarlinski/"
  description="Collection of my
personal projects, open source contribution, and a showcase of my skillsets."
  image="https://cataas.com/cat"
  imageSide="right"
/>
