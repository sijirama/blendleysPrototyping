"use client";
import { DecisionNode } from "@/lib/DecisionPlugin";
import { useCallback, useEffect, useRef, useState } from "react";

const Editor = () => {

    const _titleRef = useRef<HTMLTextAreaElement>(null);
    const ref = useRef<any>();
    const [isMounted, setIsMounted] = useState<boolean>(false);

    const initializeEditor = useCallback(async () => {
        const EditorJS = (await import("@editorjs/editorjs")).default;
        const Header = (await import("@editorjs/header")).default;
        const Embed = (await import("@editorjs/embed")).default;
        const List = (await import("@editorjs/list")).default;
        const Code = (await import("@editorjs/code")).default;
        const InlineCode = (await import("@editorjs/inline-code")).default;

        if (!ref.current) {
            const editor = new EditorJS({
                holder: "editor",
                onReady: () => {
                    ref.current = editor;
                },
                placeholder: "Type here to write your post...",
                inlineToolbar: true,
                data: { blocks: [] },
                tools: {
                    header: Header,
                    list: List,
                    code: Code,
                    inlineCode: InlineCode,
                    embed: Embed,
                    Decision:DecisionNode
                },
            });
        }
    }, []);

    useEffect(() => {
        if (typeof window !== "undefined") {
            setIsMounted(true);
        }
    }, []);

    useEffect(() => {
        const init = async () => {
            await initializeEditor();

            setTimeout(() => {
                _titleRef.current?.focus();
            }, 0);
        };
        if (isMounted) {
            init();
            return () => {
                ref.current?.destroy();
                ref.current = undefined;
            };
        }
    }, [isMounted, initializeEditor]);


    const onSubmit = async (event:React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const blocks = await ref.current?.save();

        const payload: any = {
            //title: data.title,
            title: "Mr bridgers",
            content: blocks,
            communityId: "3000",
        };
        //console.log('%o',payload);
        console.log(payload.content.blocks);
        //CreatePost(payload);
    };

    if (!isMounted) {
        return null;
    }

    return (
        <div className="w-full p-4 bg-zinc-50 rounded-lg border border-zinc-200">
            <form
                id="community-post-form"
                className="w-full"
                //onSubmit={handleSubmit(onSubmit)}
                onSubmit={(event) => onSubmit(event)}
            >
                <div className="prose prose-stone dark:prose-invert  overflow-hidden ">
                    <div id="editor" className="max-h-[300px] min-w-full flex-wrap text-wrap" />
                </div>
            </form>
        </div>
    );
};

export default Editor;













// const { mutate: CreatePost } = useMutation({
//     mutationFn: async ({ title, content, communityId }: PostType) => {
//         const payload: PostType = {
//             title,
//             content,
//             communityId,
//         };
//         const { data } = await axios.post(
//             "/api/community/post/create",
//             payload,
//         );
//         return data;
//     },
//     onError: (_err) => {
//         toast({
//             title: "Something went wrong",
//             description: "Your post was not published",
//             variant: "destructive",
//         });
//     },
//     onSuccess: () => {
//         const newpathname = pathname.split("/").slice(0, -1).join("/");
//         router.push(newpathname);
//         router.refresh();
//
//         return toast({
//             description: "Your post has been published",
//         });
//     },
// });
//
// const { ref: titleRef, ...rest } = register("title");

// useEffect(() => {
//     if (Object.keys(errors).length) {
//         for (const [_key, value] of Object.entries(errors)) {
//             toast({
//                 title: "something went wrong",
//                 description: (value as { message: string }).message,
//                 variant: "destructive",
//             });
//         }
//     }
// }, [errors]);


