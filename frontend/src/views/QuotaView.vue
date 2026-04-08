<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
import { useConsoleStore } from '@/stores/console'

const store = useConsoleStore()

function quotaPercent(used: number, total: number) {
  return Math.round((used / total) * 100)
}
</script>

<template>
  <section class="page">
    <PageHeader
      title="Quota"
      description="Use this view to visualize provider capacity and identify upcoming pressure before scheduling more tasks."
    />

    <div class="quota-grid">
      <article v-for="item in store.quotas" :key="item.provider" class="quota-card">
        <div class="quota-card__header">
          <div>
            <p class="quota-card__name">{{ item.provider }}</p>
            <p class="quota-card__time">Updated {{ item.updatedAt }}</p>
          </div>
          <p class="quota-card__percent">{{ quotaPercent(item.used, item.total) }}%</p>
        </div>
        <div class="progress progress--large">
          <div class="progress__bar" :style="{ width: `${quotaPercent(item.used, item.total)}%` }"></div>
        </div>
        <p class="quota-card__summary">{{ item.used }} GB used / {{ item.total }} GB total</p>
      </article>
    </div>
  </section>
</template>
