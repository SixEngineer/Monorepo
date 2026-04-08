<script setup lang="ts">
import { computed } from 'vue'
import type { HealthState } from '@/types/common'

const props = defineProps<{
  state: HealthState
}>()

const label = computed(() => {
  if (props.state === 'healthy') return 'Healthy'
  if (props.state === 'warning') return 'Warning'
  return 'Offline'
})
</script>

<template>
  <span class="status-badge" :class="`status-badge--${state}`">{{ label }}</span>
</template>
