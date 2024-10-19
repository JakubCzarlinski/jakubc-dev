<script lang="ts">
  import { onMount } from "svelte";

  let {
    link,
    className = "h1",
  }: {
    link: string;
    className?: string;
  } = $props();
  let copied = $state(false);

  const copyLink = (link: string) => {
    navigator.clipboard.writeText(link).then(() => {
      copied = true;
      setTimeout(() => {
        copied = false;
      }, 1000);
    });
    history.replaceState(null, "", link);
  };

  onMount(() => {
    if (link[0] === "#") link = location.origin + location.pathname + link;
  });
</script>

<button
  class="hover-text text-token relative inline w-fit text-center font-semibold"
  onclick={() => copyLink(link)}
  aria-label="Copy link"
>
  <div class="inline m-4 {className}">#</div>
  {#if copied}
    <span class="tooltip">Copied link</span>
  {/if}
</button>
