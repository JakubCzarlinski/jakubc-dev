<script lang="ts">
  import { onMount } from "svelte";

  const titles = [
    "Software Engineer",
    "ex-Google Intern x2",
    "ML Engineer",
    "Mentor",
    "Freelancer",
    "WAI President",
    "Final Year MEng",
    "Open Source Contributor",
    "Absolute Nerd",
    "Cat Enthusiast",
  ];

  let currentSubtitle = "";
  let subtitleIndex = 0;
  let charIndex = 0;
  let adding = true;

  let speed = 0.25;
  let minSpeed = 0.5;
  let maxSpeed = 5.0;
  let speedUp = true;

  function updateTitles() {
    if (speedUp) {
      if (speed < maxSpeed) {
        speed += 0.01;
      } else {
        speedUp = false;
      }
    } else {
      speed -= 0.03;

      if (speed < minSpeed) {
        speed = minSpeed;
        speedUp = true;
      }
    }

    if (adding) {
      if (charIndex < titles[subtitleIndex].length) {
        currentSubtitle += titles[subtitleIndex][charIndex];
        charIndex++;
        if (titles[subtitleIndex][charIndex] === " ") {
          setTimeout(updateTitles, 100 / speed); // Longer delay for spaces
          return;
        }
        setTimeout(updateTitles, 25 / speed); // Shorter delay for non-space characters
        return;
      }
      adding = false;
      setTimeout(updateTitles, 1000 / speed); // Pause before deleting
      return;
    }
    if (charIndex === 0) {
      adding = true;
      subtitleIndex = (subtitleIndex + 1) % titles.length;
      setTimeout(updateTitles, 25 / speed);
      return;
    }

    const words = currentSubtitle.split(" ");
    if (words.length > 1) {
      words.pop();
      currentSubtitle = words.join(" ");
    } else {
      currentSubtitle = "";
    }
    charIndex = currentSubtitle.length;
    setTimeout(updateTitles, 100 / speed); // Delay before starting to add next title
  }

  onMount(() => {
    updateTitles();
  });
</script>

<h1 class="h-0">jakubc.dev</h1>
<div
  class="hero flex h-[80svh] w-[85%] flex-row items-center justify-between p-4"
>
  <div class="mx-2 mb-2 mt-auto md:mb-24 md:ml-24">
    <h1 class="w-fit text-nowrap bg-surface-500 text-4xl">Jakub Czarliński</h1>
    <h2 class="h-8 w-fit text-nowrap bg-surface-700 text-2xl">
      {currentSubtitle}
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

<svg>
  <filter id="noiseFilter">
    <feTurbulence
      type="fractalNoise"
      baseFrequency="0.5"
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
