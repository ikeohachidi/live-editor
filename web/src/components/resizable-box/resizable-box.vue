<template>
    <div class="resizable-box-wrapper">
        <div class="resizable-box edit">
            <slot></slot>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'ResizableBox'
}
</script>

<script setup lang="ts">
</script>

<style scoped>
.resizable-box-wrapper {
    display: flex;
    flex-grow: 1;
    height: 100%;
    overflow: hidden;
    min-height: 300px;
}

.resizable-box {
    height: 100%;
    width: 300px;
    flex-grow: 1;
}

.resizer {
    width: 20px;
    height: 100%;
    cursor: col-resize;
}
</style>