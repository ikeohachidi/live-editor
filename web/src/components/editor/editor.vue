<template>
	<div class="editor-wrapper">
		<div
			class="editor"
			ref="editorEl"
			@click="onEditorBoxClick"
		>
		</div>

		<div class="bottom">
			<slot name="bottom"></slot>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { Compartment, EditorState, Extension } from '@codemirror/state';
import { EditorView, basicSetup } from "codemirror"
import { ViewUpdate, keymap } from '@codemirror/view';
import { indentWithTab } from "@codemirror/commands"
import { oneDark } from '@codemirror/theme-one-dark';

export type EditorLanguage = 'css' | 'js' | 'html' | 'sass' | 'ts';

const editorEl = ref<HTMLElement>();
let editor: EditorView;

const props = defineProps<{
	language: EditorLanguage,
	modelValue: string
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: string): void
}>();

let languages = new Compartment;
let languageCache: {[language in EditorLanguage]?: Extension} = {};

watch(() => props.language, async (lang) => {
	const langExtensions: Extension[] = [];
	const extension = languageCache[lang];

	if (extension) {
		langExtensions.push(extension);

		editor.dispatch({
			effects: languages.reconfigure(langExtensions)
		});
	}
});

const onEditorBoxClick = (): void => {
	if (editor) {
		editor.focus();
	}
}

onMounted(async() => {
	const extensions: Extension[] = [];

	if (props.language === 'js') {
		const { javascript } = await import("@codemirror/lang-javascript");
		languageCache.js = javascript();
	}

	if (props.language === 'html') {
		const { html } = await import("@codemirror/lang-html");
		languageCache.html = html();
	}

	if (props.language === 'css') {
		const { css } = await import("@codemirror/lang-css");
		const { sass } = await import("@codemirror/lang-sass");
		languageCache.css = css();
		languageCache.sass = sass();
	}

	Object.entries(languageCache)
		.forEach(([_, extension]) => {
			if (extension) {
				extensions.push(extension);
			}
		});

	let startState = EditorState.create({
		doc: props.modelValue,
		extensions: [
			basicSetup,
			languages.of(extensions),
			keymap.of([indentWithTab]),
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
.editor-wrapper {
	display: flex;
	flex-direction: column;
	/* border: 1px solid red; */
	height: 100%;
}

.editor {
	background-color: #282c34;
	flex-grow: 1;
	overflow: auto;
}

.bottom {
	margin-top: auto;
}
</style>