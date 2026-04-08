<script setup lang="ts">
import type { MetricCardData } from '@/types/dashboard'

defineProps<{
  item: MetricCardData
}>()
</script>

<template>
  <article class="metric-card">
    <p class="metric-card__title">{{ item.title }}</p>
    <p class="metric-card__value">{{ item.value }}</p>
    <p class="metric-card__detail">{{ item.detail }}</p>
    <p class="metric-card__trend">{{ item.trend }}</p>
  </article>
</template>
