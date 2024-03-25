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
    border: 1px solid green;
    overflow: hidden;
    min-height: 300px;
}

.resizable-box {
    border: 1px solid red;
    height: 100%;
    width: 300px;
    /* position: relative; */
    flex-grow: 1;
}

.resizer {
    width: 20px;
    height: 100%;
    background: blue;
    cursor: col-resize;
}
</style>