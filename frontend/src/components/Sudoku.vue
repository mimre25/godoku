<script setup lang="ts">
interface Props {
  data: Array<number>
}
const { data } = defineProps<Props>()
</script>

<template>
  <div class="flex justify-center border-red-500 border-1">
    <div class=" grid grid-cols-[2em_2em_2em_2em_2em_2em_2em_2em_2em] gaps-1">
      <div
        class="grid border-1 border-white text-center"
        :class="[idx % 3 == 0 ? 'border-l-3' : '', Math.floor(idx/9) %3 == 0? 'border-t-3' : '',
          idx > 71 ? 'border-b-3' : '', (idx +1) %9 == 0 ? 'border-r-3' : ''
        ]"
        v-for="(i, idx) in data" v-bind:key="idx"
      >
        {{ i }}
      </div>
    </div>
  </div>
</template>
