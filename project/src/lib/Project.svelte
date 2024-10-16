<script lang="ts">
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import { onMount } from "svelte";
  let {
    id = "",
    title = "",
    description = "",
    link = "",
    image = "",
    imageSide = "left",
  }: {
    id?: string;
    title?: string;
    description?: string;
    link?: string;
    image?: string;
    imageSide?: "left" | "right";
  } = $props();

  let currentDomain = $state("");
  onMount(() => {
    currentDomain = window.location.origin ?? "https://jakubc.dev";
  });
</script>

<CopyLink className={"h2"} link="{currentDomain}#{id}" />

<h2 class="wavy inline underline-offset-[12px] leading-loose">
  {#if link}
    <a href={link} target="_blank">{title}</a>
  {:else}
    {title}
  {/if}
</h2>

<div class="block mb-8"></div>

{#snippet desc(text: string)}
  <p class="text-justify align-text-top place-self-start leading-loose">
    {text}
  </p>
{/snippet}

<div class="grid grid-cols-1 md:grid-cols-2 gap-2 place-items-center">
  {#if imageSide === "right"}
    {@render desc(description)}
  {/if}

  <div class="flex aspect-square max-h-[300px] overflow-clip">
    <img
      class="aspect-square w-full place-self-center object-cover"
      src={image}
      alt={title}
    />
  </div>

  {#if imageSide === "left"}
    {@render desc(description)}
  {/if}
</div>
