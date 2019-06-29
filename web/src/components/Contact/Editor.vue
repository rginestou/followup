
<template>
  <div class="editor">
    <editor-content class="editor__content" style="height:4em" :editor="editor" />
  </div>
</template>

<script>
import { Editor, EditorContent } from 'tiptap'
import {
	Blockquote,
	BulletList,
	Heading,
	ListItem,
	OrderedList,
	TodoItem,
	TodoList,
	Bold,
	Italic,
	Link,
	Strike,
	Underline,
	History,
} from 'tiptap-extensions'
export default {
	components: {
		EditorContent,
	},
	data() {
		return {
			editor: new Editor({
				extensions: [
					new Blockquote(),
					new BulletList(),
					new Heading({ levels: [1, 2, 3] }),
					new ListItem(),
					new OrderedList(),
					new TodoItem(),
					new TodoList(),
					new Link(),
					new Bold(),
					new Italic(),
					new Strike(),
					new Underline(),
					new History(),
				],
				content: `Contenu...`,
				onUpdate: ({ getHTML }) => {
					// const newContent = this.getHTML()
				},
			}),
		}
	},
	beforeDestroy() {
		this.editor.destroy()
	},
}
</script>
