<script lang="ts">
  import Blog from "@/project/src/lib/Blog.svelte";
  import Hero from "@/project/src/lib/Hero.svelte";
  import Portfolio from "@/project/src/lib/Portfolio.svelte";
  import Top from "@/project/src/lib/Top.svelte";
</script>

<div id="top" class="h-0"></div>
<div class="h-[10svh]"></div>
<div class="grid grid-cols-1 place-items-center">
  <Hero />
</div>

<div class="grid grid-cols-1 place-items-center px-4">
  <div class="w-full sm:w-[80%] md:w-[70%]">
    <div class="italic h-[5svh]">
      Hey~
    </div>
    <div class="italic font-bold h-[5svh]">
      Nice to see you here so early! This page is still under construction, but
      feel free to look around.
    </div>
    <div class="italic h-[5svh]">
      - Jakub
    </div>

    <div id="portfolio">
      <Portfolio />
    </div>

    <div class="h-[10vh]"></div>

    <div id="blog">
      <Blog />
      <Top />
    </div>

    <div class="h-[10vh]"></div>
  </div>
</div>
