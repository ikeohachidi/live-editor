<template>
	<div
		class="editor"
		ref="editorEl"
		@click="onEditorBoxClick"
	>
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { EditorState } from '@codemirror/state';
import { EditorView, basicSetup } from "codemirror"
import { ViewUpdate } from '@codemirror/view';
import { oneDark } from '@codemirror/theme-one-dark';

export type EditorLanguage = 'css' | 'javascript' | 'html';

const editorEl = ref<HTMLElement>();
let editor: EditorView;

const props = defineProps<{
	language: EditorLanguage,
	modelValue: string
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: string): void
}>();

const onEditorBoxClick = (): void => {
	if (editor) {
		editor.focus();
	}
}

onMounted(async() => {
	const extensions = [basicSetup];

	if (props.language === 'javascript') {
		const { javascript } = await import("@codemirror/lang-javascript");
		extensions.push(javascript());
	}

	if (props.language === 'html') {
		const { html } = await import("@codemirror/lang-html");
		extensions.push(html());
	}

	if (props.language === 'css') {
		const { css } = await import("@codemirror/lang-css");
		extensions.push(css());
	}


	let startState = EditorState.create({
		doc: props.modelValue,
		extensions: [
			...extensions,
			EditorView.updateListener.of((view: ViewUpdate) => {
				if (view.docChanged) {
					const content = editor.state.doc.toString()
					emit('update:modelValue', content);
				}
			}),
			oneDark
		]
	})

	editor = new EditorView({
		state: startState,
		parent: editorEl.value
	});
})
</script>

<style scoped>
.editor {
	height: 100%;
	background-color: #282c34;
}
</style>