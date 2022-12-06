import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PostCreateDialogComponent} from "./post-create-dialog.component";

const routes: Routes = [{path: '', component: PostCreateDialogComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PostCreateDialogRoutingModule { }
