<template>
  <div class="bg-section slidev-layout new-section">
    <CornerCurves class="absolute bottom-0 right-0 transform scale-x--100" />
    <CornerCurves class="absolute left-0 top-0 transform rotate-90" />
    <div class="flex flex-col justify-center text-center">
      <slot></slot>
    </div>
  </div>
</template>

<style>
.slidev-layout.new-section {
  h1 {
    @apply text-primary text-4xl leading-20 mb-0 font-title font-extrabold;
  }

  h1 + p {
    @apply mt-4 opacity-100;
  }

  p > img {
    @apply w-2/3 mx-auto;
  }
}
</style>
<script setup lang="ts">
</script>
