<script setup lang="ts">
defineProps<{
  title: string
  description: string
}>()
</script>

<template>
  <div class="page-header">
    <div>
      <p class="page-header__eyebrow">Module</p>
      <h2>{{ title }}</h2>
    </div>
    <p class="page-header__description">{{ description }}</p>
  </div>
</template>
