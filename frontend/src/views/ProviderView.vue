<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
import StatusBadge from '@/components/common/StatusBadge.vue'
import { useConsoleStore } from '@/stores/console'

const store = useConsoleStore()
</script>

<template>
  <section class="page">
    <PageHeader
      title="Providers"
      description="A management view for adapter state, capability display, and future secondary-auth configuration."
    />

    <div class="provider-grid">
      <article v-for="provider in store.providers" :key="provider.id" class="provider-card">
        <div class="provider-card__header">
          <div>
            <p class="provider-card__name">{{ provider.name }}</p>
            <p class="provider-card__id">{{ provider.id }}</p>
          </div>
          <StatusBadge :state="provider.status" />
        </div>
        <p class="provider-card__section-title">Capability</p>
        <p class="provider-card__text">{{ provider.capability }}</p>
        <p class="provider-card__section-title">Auth Mode</p>
        <p class="provider-card__text">{{ provider.authMode }}</p>
      </article>
    </div>
  </section>
</template>
