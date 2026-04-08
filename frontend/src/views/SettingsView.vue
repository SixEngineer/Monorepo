<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
</script>

<template>
  <section class="page">
    <PageHeader
      title="Settings"
      description="Group future OpenList, aria2, download-policy, quota-policy, and user-preference forms here."
    />

    <div class="placeholder-grid">
      <article class="placeholder-card">
        <h3>OpenList Settings</h3>
        <p>Endpoint, credentials, and adapter defaults.</p>
      </article>
      <article class="placeholder-card">
        <h3>aria2 Settings</h3>
        <p>RPC host, concurrency, and scheduling behavior.</p>
      </article>
      <article class="placeholder-card">
        <h3>Download Policy</h3>
        <p>Path defaults, retry rules, and fallback behavior.</p>
      </article>
    </div>
  </section>
</template>
