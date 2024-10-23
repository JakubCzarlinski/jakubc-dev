<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import Project from "@/project/src/lib/Project.svelte";
</script>

<CopyLink link="#portfolio" />

<h1 class="wavy mb-12 inline-block underline-offset-[12px]">Portfolio</h1>

<hr class="!border-t-2" />

<Project
  id="framework"
  title="Go-Svelte Framework"
  description="A blazingly
fast full-stack framework that Server-Side-Renders Svelte components into Go
templates at compile time. It combines the full expressivity of Svelte with the
speed of Go, creating an ideal synergy. This single binary is designed to
minimize memory footprint while preserving the rapid development experience of
the JavaScript ecosystem. It's a dream come true for developers building fast,
scalable, and visually appealing web applications."
  image="https://raw.githubusercontent.com/JakubCzarlinski/jakubc-dev/refs/heads/main/public/ginger.svg"
  imageCaption="Caesar confused after being woken up."
  initialOpen
>
  {#snippet chips()}
    <div class="flex flex-wrap">
      <img
        src="https://img.shields.io/badge/Bun-%23000000.svg?style=for-the-badge&logo=bun&logoColor=white&cacheSeconds=604800"
        alt="Bun"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Alpine-%230D597F?style=for-the-badge&logo=alpine-linux&logoColor=white&cacheSeconds=604800"
        alt="Alpine Linux"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white&cacheSeconds=604800"
        alt="PostgreSQL"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white&cacheSeconds=604800"
        alt="Vite"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white&cacheSeconds=604800"
        alt="TypeScript"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/gin-%23008ECF?style=for-the-badge&logo=gin&logoColor=white&cacheSeconds=604800"
        alt="Gin"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white&cacheSeconds=604800"
        alt="Go"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white&cacheSeconds=604800"
        alt="Docker"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white&cacheSeconds=604800"
        alt="Tailwind CSS"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white&cacheSeconds=604800"
        alt="Cloudflare"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white&cacheSeconds=604800"
        alt="Svelte"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/RollupJS-ef3335?style=for-the-badge&logo=rollup.js&logoColor=white&cacheSeconds=604800"
        alt="RollupJS"
        height="28"
      />
    </div>
  {/snippet}
  {#snippet bottonSnippet()}
    <div class="text-justify leading-loose space-y-4">
      <h3 class="italic font-bold wavy">
        Marketing fluff aside, what do the numbers say?
      </h3>
      <p class="indent-4 pl:8 md:pl-16">
        Minimal build time for a full-stack app is just
        <span class="wavy">0.6 seconds.</span>
        Running this website inside an Alpine Docker container, the framework handles
        over <span class="wavy">400K requests per second.</span>
        For comparison, the `<code>hello world</code>` example for Bun's server
        can handle about 100k requests per second in the same environment. That
        is <span class="wavy">3.26GB versus 11.55MB</span> of data transferred.
      </p>

      <h3 class="italic font-bold wavy">So what about memory footprint?</h3>
      <p class="indent-4 pl:8 md:pl-16">
        Bun uses 42.8MB of memory on standby and peaks at 140MB during
        benchmarking. 42.8MB for a `<code>hello world</code>` doing nothing?!?
        Go-Svelte uses
        <span class="wavy">just 5.0MB and 40.1MB respectively.</span>
      </p>

      <h3 class="italic font-bold wavy">
        What's the big deal? Aren't computers fast enough?
      </h3>

      <p class="indent-4 pl:8 md:pl-16">
        Sure, modern computers are fast. Going faster means you can do more with
        less. Less memory, less power, less cost. It's about efficiency. When
        your whole docker container can fit into the L3 cache of a modern
        desktop CPU, and the application
        <span class="wavy">fits in L2 cache,</span> you know you're doing
        <span class="italic">something</span> right. Fewer cache misses, fewer context
        switches, fewer bottlenecks. It's about making the most of what you have.
        There is no need to waste resources, bloating your server and client. Have
        you not been annoyed by sluggish software in the last 5 years?
      </p>

      <p class="indent-4 pl:8 md:pl-16">
        Just because we can use gigabytes of memory, a
        <code>node_modules</code> folder large enough to be a fill a data centre,
        and a CPU running hotter than a nuclear reactor, doesn't mean we should.
        It's about doing more with less. It's about being efficient.
      </p>
    </div>
  {/snippet}
</Project>

<Project
  id="google"
  title="Google Internships"
  description="Google STEP intern
2022, working on Android Google Search. Google SWE intern 2023, working on
Text-to-Speech research. Specifics omitted."
  image="https://raw.githubusercontent.com/JakubCzarlinski/jakubc-dev/refs/heads/main/public/caesar_laying.svg"
  imageCaption="Caesar basking in the sun next to a plant."
  imageSide="right"
>
  {#snippet chips()}
    <div class="flex flex-wrap">
      <img
        src="https://img.shields.io/badge/numpy-%23013243.svg?style=for-the-badge&logo=numpy&logoColor=white&cacheSeconds=604800"
        alt="NumPy"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54&cacheSeconds=604800"
        alt="Python"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white&cacheSeconds=604800"
        alt="Go"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Jupyter-F37626?style=for-the-badge&logo=Jupyter&logoColor=white&cacheSeconds=604800"
        alt="Jupyter"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/TensorFlow-%23FF6F00.svg?style=for-the-badge&logo=TensorFlow&logoColor=white&cacheSeconds=604800"
        alt="TensorFlow"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/JAX-%23ab00b6?style=for-the-badge&logoColor=white&cacheSeconds=604800"
        alt="JAX"
        height="28"
      />
    </div>
  {/snippet}
</Project>

<Project
  id="dissertation"
  title="Diffusion Multi-Speaker TTS Synthesis"
  description="Third-year dissertation project focused on applying diffusion
models to generate speech from multiple speakers. Inspired by the NaturalSpeech2
model developed by researchers at Microsoft, but iterated on by conditioning the
model with speaker embeddings rather than whole sections of audio."
  image="https://raw.githubusercontent.com/JakubCzarlinski/jakubc-dev/refs/heads/main/public/black_cat.svg"
  imageCaption="A neighbourhood cat watching me clean the garden."
>
  {#snippet chips()}
    <div class="flex flex-wrap">
      <img
        src="https://img.shields.io/badge/numpy-%23013243.svg?style=for-the-badge&logo=numpy&logoColor=white&cacheSeconds=604800"
        alt="NumPy"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Overleaf-%23336699.svg?style=for-the-badge&logo=overleaf&logoColor=white&cacheSeconds=604800"
        alt="Overleaf"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54&cacheSeconds=604800"
        alt="Python"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/PyTorch-%23EE4C2C.svg?style=for-the-badge&logo=PyTorch&logoColor=white&cacheSeconds=604800"
        alt="PyTorch"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Jupyter-F37626?style=for-the-badge&logo=Jupyter&logoColor=white&cacheSeconds=604800"
        alt="Jupyter"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/Hugging%20Face-%23FFD21E?style=for-the-badge&logo=Hugging-Face&logoColor=black&cacheSeconds=604800"
        alt="Hugging Face"
        height="28"
      />
      <img
        src="https://img.shields.io/badge/conda-342B029.svg?&style=for-the-badge&logo=anaconda&logoColor=white&cacheSeconds=604800"
        alt="Conda"
        height="28"
      />
    </div>
  {/snippet}
</Project>

<Project
  id="github"
  title="GitHub"
  description="Collection of many of my personal
projects, open-source contribution, and a showcase of my skillsets. Contains a
variety of projects, from full-stack web applications, to machine learning
research projects, to small utility scripts. It's a great place to see what I've
been up to. Whenever I have a new idea, I try to make it a reality and share it
with the world."
  image="https://raw.githubusercontent.com/JakubCzarlinski/jakubc-dev/refs/heads/main/public/luna.svg"
  imageCaption="Luna; a very friendly and free-spirited cat."
  imageSide="right"
>
  {#snippet bottonSnippet()}
    <div class="text-justify leading-loose space-y-4">
      <h3 class="italic font-bold wavy">
        Isn't that a little too many badges? What do they even mean?
      </h3>
      <p class="indent-4 pl:8 md:pl-16">
        Sure is! Feel free to doubt my experience, but I can assure you that I
        have interacted with all of these technologies. It's always fun to try a
        new technology, program using a new paradigm, learn something new. It's
        what makes the field of Computer Science so exciting; ever-changing. It
        feels very enabling to be able to pick up any tool and just start
        building. Given that I started programming very early, I've had the
        opportunity to explore my interests and opportunities.
      </p>
    </div>
    <a
      class="wavy hover-text text-token pt-8"
      href="https://www.github.com/JakubCzarlinski/"
      target="_blank"
    >
      https://www.github.com/JakubCzarlinski/
    </a>
  {/snippet}
</Project>
