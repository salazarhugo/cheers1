import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PartyTicketCreateComponent} from "./party-ticket-create.component";


const routes: Routes = [{path: '', component: PartyTicketCreateComponent}];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class PartyTicketCreateRoutingModule {
}
