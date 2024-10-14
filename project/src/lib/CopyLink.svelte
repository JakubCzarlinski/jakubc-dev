<script lang="ts">
  let { link }: { link: string } = $props();
  let copied = $state(false);

  const copyLink = (link: string) => {
    navigator.clipboard.writeText(link).then(() => {
      copied = true;
      setTimeout(() => {
        copied = false;
      }, 1000);
    });
  };
</script>

<button
  class="hover-text text-token inline w-fit text-center font-semibold relative"
  onclick={() => copyLink(link)}
  aria-label="Copy link"
>
  <h1 class="inline">#</h1>
  {#if copied}
    <span class="tooltip">Copied link</span>
  {/if}
</button>
