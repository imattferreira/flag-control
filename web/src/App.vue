<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <RouterView v-slot="{ Component }">
    <template v-if="Component">
      <Transition mode="out-in">
        <KeepAlive>
          <Suspense>
            <component :is="Component"></component>
            <template #fallback>Loading...</template>
          </Suspense>
        </KeepAlive>
      </Transition>
    </template>
  </RouterView>
</template>
