<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
import { useConsoleStore } from '@/stores/console'

const store = useConsoleStore()
</script>

<template>
  <section class="page">
    <PageHeader
      title="Download Tasks"
      description="A task console for queue visibility, state transitions, progress, and later task actions such as retry or cancel."
    />

    <section class="panel">
      <div class="table">
        <div class="table__row table__row--head">
          <span>Task</span>
          <span>Provider</span>
          <span>Target</span>
          <span>Speed</span>
          <span>Progress</span>
          <span>Status</span>
        </div>
        <div v-for="task in store.taskTable" :key="task.id" class="table__row">
          <span>{{ task.name }}</span>
          <span>{{ task.provider }}</span>
          <span>{{ task.target }}</span>
          <span>{{ task.speed }}</span>
          <span>{{ task.progress }}%</span>
          <span class="table__status">{{ task.status }}</span>
        </div>
      </div>
    </section>
  </section>
</template>
