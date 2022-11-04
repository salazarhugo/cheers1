import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

const routes: Routes = [
    {
        path: 'feed',
        loadChildren: () => import('../party-feed/party-feed.module').then(m => m.PartyFeedModule)
    },
    {
        path: 'create',
        loadChildren: () => import('../party-create/party-create.module').then(m => m.PartyCreateModule)
    },
    {
        path: ':id/edit',
        loadChildren: () => import('../party-edit/party-edit.module').then(m => m.PartyEditModule)
    },
    {
        path: ':id/guestlist',
        loadChildren: () => import('../party-guest-list/party-guest-list.module').then(m => m.PartyGuestListModule)
    },
    {
        path: ':id/tickets',
        loadChildren: () => import('../party-tickets/party-tickets.module').then(m => m.PartyTicketsModule)
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class PartyShellRoutingModule {
}
