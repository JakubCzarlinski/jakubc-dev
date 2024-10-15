<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import { onMount } from "svelte";
  let {
    id = "",
    title = "",
    description = "",
    link = "",
    image = "",
  }: {
    id?: string;
    title?: string;
    description?: string;
    link?: string;
    image?: string;
  } = $props();

  let currentDomain = $state("");
  onMount(() => {
    currentDomain = window.location.origin ?? "https://jakubc.dev";
  });
</script>

<CopyLink className={"h2"} link="{currentDomain}#{id}" />

<h2 class="wavy inline underline-offset-[12px] leading-loose">
  {#if link}
    <a href={link} target="_blank">{title.toLowerCase()}</a>
  {:else}
    {title.toLowerCase()}
  {/if}
</h2>

<div class="block mb-8"></div>

<div class="grid grid-cols-1 md:grid-cols-2 gap-2 place-items-center">
  <div class="flex aspect-square max-h-[300px] overflow-clip">
    <img
      class="aspect-square w-full place-self-center object-cover"
      src={image}
      alt={title}
    />
  </div>

  <p class="lowercase text-justify">
    {description}
  </p>
</div>
