<script setup lang="ts">
import { ref } from 'vue'
import Sudoku from './components/Sudoku.vue'

enum MsgType {
  INIT = 0,
  CONTROL = 1,
  UPDATE = 2,
}

enum ControlMsgType {
  START_SOLVING = 0,
}
const socket = new WebSocket('ws://localhost:8123/ws')

socket.addEventListener('open', (event) => {
  console.log('Message from server ', event)
})

const sudoku = ref(new Uint8Array([]))

const handleWsMsg = (msgType: MsgType, ...data: Uint8Array) => {
  console.log(`Got message of type ${msgType}`)
  switch (msgType) {
    case MsgType.UPDATE:
      sudoku.value = data
      console.log('sudoku', sudoku)
      break
  }
}

socket.addEventListener('message', (event) => {
  console.log('received: ', event)
  event.data.bytes().then((b: Uint8Array) => handleWsMsg(...b))
})

const running = ref(false)

const startSolving = () => {
  if (!running.value) {
    console.log('start')
    running.value = true
    socket.send(new Uint8Array([MsgType.CONTROL, ControlMsgType.START_SOLVING]))
  }
}
</script>

<template>
  <Sudoku :data="sudoku" />
  <br />
  <div class="flex justify-center border-green-500 border-1">
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      :class="[running ? 'opacity-50 cursor-not-allowed' : '']"
      :disabled="running"
      :onclick="startSolving"
    >
      Start
    </button>
  </div>
</template>

<style scoped></style>
