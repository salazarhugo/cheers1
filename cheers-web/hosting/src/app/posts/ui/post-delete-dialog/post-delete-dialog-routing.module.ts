import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PostDeleteDialogComponent} from "./post-delete-dialog.component";

const routes: Routes = [{path: '', component: PostDeleteDialogComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PostDeleteDialogRoutingModule { }
