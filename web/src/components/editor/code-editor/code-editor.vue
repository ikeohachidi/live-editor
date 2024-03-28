<template>
    <editor :language="codeLanguage" v-model="codeContent">
        <template #bottom>
            <div class="language">
                <button
                    v-for="l in languageList"
                    :key="l.value"
                    @click="codeLanguage = l.value"
                >
                    {{ l.label }}
                </button>
            </div>
        </template>
    </editor>
</template>

<script setup lang="ts">
import { computed, watch } from 'vue';
import Editor, { type EditorLanguage } from '../../editor/editor.vue';
import { runUpdate } from '../../../http';

const props = defineProps<{
    sessionId: number,
    content: string,
    language: EditorLanguage,
    languageList?: {label: string, value: EditorLanguage}[]
}>();

const emits = defineEmits<{
    (e: 'update:content', value: string): void,
    (e: 'update:language', value: EditorLanguage): void
}>();

const codeContent = computed({
    get() {
        return props.content;
    },
    set(value: string) {
        emits('update:content', value);
    }
});

const codeLanguage = computed({
    get() {
        return props.language;
    },
    set(value: EditorLanguage) {
        emits('update:language', value);
    }
});

watch(codeContent, (content: string) => {
	runUpdate(props.sessionId, { content, lang: props.language });
});
</script>

<style scoped>
.language button {
	margin: 5px 5px;
}

.language button:last-of-type {
    margin-right: 0;
}
</style>

