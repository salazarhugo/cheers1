import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {StoriesComponent} from "./features/story-feed/stories.component";
import {MaterialModule} from "../material/material.module";
import {MatIconModule} from "@angular/material/icon";
import {PostsModule} from "../posts/posts.module";


const components = [
    StoriesComponent,
]

@NgModule({
    declarations: [
        ...components,
    ],
    imports: [
        CommonModule,
        MaterialModule,
        MatIconModule,
    ],
    exports: [
        ...components,
    ]
})
export class StoriesModule {
}
