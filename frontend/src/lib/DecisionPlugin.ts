//@ts-nochec
import React from 'react'
import ReactDOM from 'react-dom';
//import EditorJS, { API, ToolSettings, BaseTool } from '@editorjs/editorjs';
import type { BaseTool } from '@editorjs/editorjs';

//
// type SimpleImageConstructorParams = {
//   api: API
//   config?: ToolSettings
//   data: SimpleImageData
// }

export class DecisionNode {

    static get toolbox() {
        return {
            title: 'Decision',
            icon: `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="#800080">
      <path d="M12 2l-6 6 6 6 6-6-6-6z"/>
    </svg>`
        };
    }


    render() {
        //The render method will create a UI of a Block that will be appended when our Tool will be selected from the Toolbox.
        //return document.createElement('input');
         const wrapper = document.createElement('div');
        wrapper.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="#800080"><path d="M12 2l-6 6 6 6 6-6-6-6z"/></svg>';
        return wrapper;
        
    }

    save(blockContent: any) {
        return {
            url: blockContent.value
        }
    }
}
