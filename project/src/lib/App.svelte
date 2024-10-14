<script lang="ts">
  import Blog from "@/project/src/lib/Blog.svelte";
  import Portfolio from "@/project/src/lib/Portfolio.svelte";
  import Top from "@/project/src/lib/Top.svelte";
</script>
<div id="top" class="h-0"></div>
<div class="h-[10vh]"></div>
<div class="grid grid-cols-1 place-items-center">
  <h1 class="h-0">jakubc.dev</h1>
  <div
    class="hero flex h-[80vh] w-[85%] flex-row items-center justify-between p-4"
  >
    <div class="mx-2 mb-2 mt-auto md:mb-24 md:ml-24">
      <h1 class="text-nowrap bg-surface-500 text-4xl">Jakub Czarliński</h1>
      <h2 class="w-fit text-nowrap bg-surface-700 text-2xl">
        Software Engineer
      </h2>
      <div class="mt-24 gap-4">
        <a
          class="variant-filled-surface text-token inline w-fit p-2 text-center font-semibold"
          href="#portfolio">portfolio</a
        >
        <a
          class="variant-filled-surface text-token inline w-fit p-2 text-center font-semibold"
          href="#blog">blog</a
        >
      </div>
    </div>
    <div class="blob-cont relative flex flex-col items-center justify-center">
      <div class="yellow blob"></div>
      <div class="red blob"></div>
      <div class="green blob"></div>
    </div>
  </div>
</div>
<div class="h-[20vh]"></div>

<div class="grid grid-cols-1 place-items-center px-4">
  <div class="w-full md:w-[60%]">
    <div id="portfolio">
      <Portfolio /><Top />
    </div>

    <div class="h-[10vh]"></div>

    <div id="blog">
      <Blog />
      <Top />
    </div>

    <div class="h-[10vh]"></div>

  </div>
</div>

<svg>
  <filter id="noiseFilter">
    <feTurbulence
      type="fractalNoise"
      baseFrequency="0.6"
      stitchTiles="stitch"
    />
    <feColorMatrix
      in="colorNoise"
      type="matrix"
      values="1 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 1 0"
    />
    <feComposite operator="in" in2="SourceGraphic" result="monoNoise" />
    <feBlend in="SourceGraphic" in2="monoNoise" mode="screen" />
  </filter>
</svg>
