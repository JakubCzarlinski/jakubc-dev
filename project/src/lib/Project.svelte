<script lang="ts">
  import Accordion from "@/project/src/lib/Accordion.svelte";
  import CopyLink from "@/project/src/lib/CopyLink.svelte";
  import Top from "@/project/src/lib/Top.svelte";
  import { type Snippet } from "svelte";
  let {
    id,
    title = "",
    image = "",
    imageCaption = "",
    imageSide = "left",
    initialOpen = false,
    description,
    chips,
    bottonSnippet,
  }: {
    id: string;
    title?: string;
    image?: string;
    imageCaption?: string;
    imageSide?: "left" | "right";
    initialOpen?: boolean;
    description?: Snippet | string;
    chips?: Snippet;
    bottonSnippet?: Snippet;
  } = $props();

  const gridDirection = imageSide === "left" ? "order-last" : "order-first";
</script>

<svelte:head>
  <link rel="image" href={image} />
</svelte:head>

<Accordion {initialOpen} {id}>
  {#snippet head()}
    <div>
      <CopyLink className={"h2"} link="#{id}" />
      <h2 class="wavy inline underline-offset-[12px] leading-loose">
        {title}
      </h2>
    </div>
  {/snippet}

  {#snippet details()}
    {@render chips?.()}
    <div class="block mb-8"></div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-2 place-items-center">
      <div class="grid grid-cols-1 place-items-center {gridDirection}">
        <div
          class="flex aspect-square min-h-[300px] max-h-[300px] overflow-clip"
        >
          <img
            class="aspect-square w-full place-self-center object-cover"
            src={image}
            alt={title}
          />
        </div>
        <h3 class="text-justify italic">
          {imageCaption}
        </h3>
      </div>

      <p class="text-justify align-text-top place-self-start leading-loose">
        {#if typeof description === "function"}
          {@render description()}
        {:else}
          {description}
        {/if}
      </p>
    </div>

    <div class="mt-4 grid grid-cols-1 place-items-center">
      {@render bottonSnippet?.()}
    </div>

    <Top />
  {/snippet}
</Accordion>
