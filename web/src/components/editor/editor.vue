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
import { ref, onMounted, watch } from 'vue';
import { Compartment, EditorState, Extension } from '@codemirror/state';
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

const changeLanguage = (language: EditorLanguage): void => {
	emit('update:language', language);
}

onMounted(async() => {
	const extensions: Extension[] = [];

	if (props.language === 'javascript') {
		const { javascript } = await import("@codemirror/lang-javascript");
		languageCache.javascript = javascript();
	}

	if (props.language === 'html') {
		const { html } = await import("@codemirror/lang-html");
		languageCache.javascript = html();
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