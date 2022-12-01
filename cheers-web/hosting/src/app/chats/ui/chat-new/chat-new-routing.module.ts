import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ChatNewComponent} from "./chat-new.component";

const routes: Routes = [{path: '', component: ChatNewComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChatNewRoutingModule { }
