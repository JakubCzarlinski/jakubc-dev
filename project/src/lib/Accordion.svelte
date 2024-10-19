<script lang="ts">
  import type { Snippet } from "svelte";
  import { expoInOut } from "svelte/easing";
  import { slide } from "svelte/transition";

  let {
    id,
    initialOpen = false,
    head,
    details,
  }: {
    id?: string;
    initialOpen?: boolean;
    head?: Snippet;
    details?: Snippet;
  } = $props();

  let open = $state(initialOpen);
  const onclick = () => (open = !open);
</script>

<div class="my-4" {id}>
  <!-- svelte-ignore a11y_invalid_attribute -->
  <a
    class="flex w-full justify-between items-center"
    {onclick}
    aria-expanded={open}
    aria-controls="accordion"
    href="javascript:void(0)"
  >
    {@render head?.()}
    <div class="mr-4 font-extrabold">V</div>
  </a>
  {#if open}
    <div
      class="p-4"
      transition:slide={{
        duration: 300,
        easing: expoInOut,
        axis: "y",
      }}
    >
      {@render details?.()}
    </div>
  {/if}
</div>

<hr class="!border-t-2" />
