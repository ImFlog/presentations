<script setup lang="ts">
</script>

<template>
  <div class="slidev-layout">
    <CornerCurves class="absolute bottom-0 right-0 transform scale-x--100" />
    <div
      class="grid grid-cols-2"
      :class="$attrs.reverse ? 'gap-16' : 'gap-8'"
    >
      <div
        class="prose pr-16"
        :class="{ 'order-1': $attrs.reverse }"
      >
        <slot name="default"></slot>
      </div>
      <figure class="relative flex flex-col justify-center">
        <img
          :src="$attrs.media"
          class="rounded-lg shadow-lg object-cover z-10"
        />
        <figcaption class="mt-2 text-xs w-full">
          {{ $attrs.caption }}
        </figcaption>
      </figure>
    </div>
  </div>
</template>
