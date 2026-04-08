<script setup lang="ts">
import MetricCard from '@/components/common/MetricCard.vue'
import PageHeader from '@/components/common/PageHeader.vue'
import StatusBadge from '@/components/common/StatusBadge.vue'
import { useConsoleStore } from '@/stores/console'

const store = useConsoleStore()
</script>

<template>
  <section class="page">
    <PageHeader
      title="Dashboard"
      description="A first-pass overview of service health, task flow, and quota pressure across the OpenBridge platform."
    />

    <div class="grid grid--metrics">
      <MetricCard v-for="item in store.metrics" :key="item.title" :item="item" />
    </div>

    <div class="dashboard-panels">
      <section class="panel">
        <div class="panel__header">
          <h3>System Health</h3>
          <p>{{ store.healthyServices }}/{{ store.statuses.length }} services healthy</p>
        </div>
        <div class="status-list">
          <article v-for="item in store.statuses" :key="item.name" class="status-row">
            <div>
              <p class="status-row__name">{{ item.name }}</p>
              <p class="status-row__detail">{{ item.detail }}</p>
            </div>
            <StatusBadge :state="item.state" />
          </article>
        </div>
      </section>

      <section class="panel">
        <div class="panel__header">
          <h3>Recent Alerts</h3>
          <p>Signals worth surfacing during demos and later backend integration.</p>
        </div>
        <div class="alert-list">
          <article
            v-for="item in store.alerts"
            :key="item.title"
            class="alert-card"
            :class="`alert-card--${item.level}`"
          >
            <p class="alert-card__title">{{ item.title }}</p>
            <p class="alert-card__detail">{{ item.detail }}</p>
          </article>
        </div>
      </section>
    </div>

    <section class="panel">
      <div class="panel__header">
        <h3>Recent Tasks</h3>
        <p>Download orchestration at a glance.</p>
      </div>
      <div class="task-digest-list">
        <article v-for="task in store.tasks" :key="task.id" class="task-digest">
          <div>
            <p class="task-digest__name">{{ task.name }}</p>
            <p class="task-digest__meta">{{ task.id }} · {{ task.provider }}</p>
          </div>
          <div class="task-digest__right">
            <span class="task-digest__status">{{ task.status }}</span>
            <div class="progress">
              <div class="progress__bar" :style="{ width: `${task.progress}%` }"></div>
            </div>
            <span class="task-digest__progress">{{ task.progress }}%</span>
          </div>
        </article>
      </div>
    </section>
  </section>
</template>
