<script setup lang="ts">
import Sudoku from './components/Sudoku.vue'
</script>

<template>
  <Sudoku
    :data="[1, 2, 3, 4, 5, 6, 7, 8, 9,
      1, 2, 3, 0, 0, 0, 0, 0, 0,
      2, 2, 3, 0, 0, 0, 0, 0, 0,
      3, 2, 3, 0, 0, 0, 0, 0, 0,
      4, 2, 3, 0, 0, 0, 0, 0, 0,
      5, 2, 3, 0, 0, 0, 0, 0, 0,
      6, 2, 3, 0, 0, 0, 0, 0, 0,
      7, 2, 3, 0, 0, 0, 0, 0, 0,
      8, 2, 3, 0, 0, 0, 0, 0, 0]
    "
  />
</template>

<style scoped></style>
