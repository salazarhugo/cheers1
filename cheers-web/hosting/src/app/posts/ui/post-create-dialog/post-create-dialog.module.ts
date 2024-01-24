import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PostCreateDialogRoutingModule} from './post-create-dialog-routing.module';
import {PostCreateDialogComponent} from './post-create-dialog.component';
import {MatButtonModule} from "@angular/material/button";
import {NgxDropzoneModule} from "ngx-dropzone";


@NgModule({
    declarations: [
        PostCreateDialogComponent
    ],
    imports: [
        CommonModule,
        PostCreateDialogRoutingModule,
        MatButtonModule,
        NgxDropzoneModule
    ],
    exports: [
        PostCreateDialogComponent
    ],
})
export class PostCreateDialogModule {
}
