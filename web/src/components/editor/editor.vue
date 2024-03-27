<template>
	<div>
		<button @click="changeLanguage('sass')" v-if="language === 'css'">Sass</button>
		<button v-else-if="language === 'sass'" @click="changeLanguage('css')">Css</button>

		<div
			class="editor"
			ref="editorEl"
			@click="onEditorBoxClick"
		>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { EditorState } from '@codemirror/state';
import { EditorView, basicSetup } from "codemirror"
import { ViewUpdate } from '@codemirror/view';
import { oneDark } from '@codemirror/theme-one-dark';

export type EditorLanguage = 'css' | 'javascript' | 'html' | 'sass';

const editorEl = ref<HTMLElement>();
let editor: EditorView;

const props = defineProps<{
	language: EditorLanguage,
	modelValue: string
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: string): void
	(e: 'update:language', value: EditorLanguage): void
}>();

const onEditorBoxClick = (): void => {
	if (editor) {
		editor.focus();
	}
}

// const loadedLanguages: {[language in EditorLanguage]?: Extension} = {};

// watch(() => props.language, (lang: EditorLanguage) => {
// 	if (lang === 'sass') {
// 		if (loadedLanguages.sass) {
// 			editor.
// 		}
// 	}
// });

const changeLanguage = (language: EditorLanguage): void => {
	emit('update:language', language);
}

onMounted(async() => {
	const extensions = [basicSetup];

	if (props.language === 'javascript') {
		const { javascript } = await import("@codemirror/lang-javascript");
		extensions.push(javascript());
	}

	if (props.language === 'html') {
		const { html } = await import("@codemirror/lang-html");
		// loadedLanguages.html = html();
		extensions.push(html());
	}

	if (props.language === 'css') {
		const { css } = await import("@codemirror/lang-css");
		// const { sass } = await import("@codemirror/lang-sass");
		// loadedLanguages.css = css();
		// loadedLanguages.sass = sass();
		extensions.push(css());
	}

	// Object.entries(loadedLanguages)
	// 	.forEach(([_, extension]) => {
	// 		if (extension) {
	// 			extensions.push(extension);
	// 		}
	// 	})


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