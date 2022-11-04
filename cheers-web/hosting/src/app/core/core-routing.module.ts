import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ChatComponent} from "../chats/feature/chat.component";
import {HomeComponent} from "../home/features/home/home.component";
import {MapComponent} from "../shared/ui/map/map.component";
import {TicketsComponent} from "../shared/ui/tickets/tickets.component";

const routes: Routes = [
    {
        path: 'home',
        component: HomeComponent,
    },
    {
        path: 'tickets',
        component: TicketsComponent,
    },
    {
        path: 'friends',
        component: ChatComponent,
    },
    {
        path: 'events',
        redirectTo: 'parties',
    },
    {
        path: 'map',
        component: MapComponent,
    },
    {
        path: 'notifications',
        loadChildren: () => import('../activities/feature/activity-shell/activity-shell.module').then(m => m.ActivityShellModule)
    },
    {
        path: 'chats',
        component: ChatComponent,
        loadChildren: () => import('../chats/chats.module').then(m => m.ChatsModule)
    },
    {
        path: 'parties',
        loadChildren: () => import('../parties/feature/party-shell/party-shell.module').then(m => m.PartyShellModule)
    },
    {
        path: 'settings',
        loadChildren: () => import('../settings/feature/setting-shell/setting-shell.module').then(m => m.SettingShellModule)
    },
    {
        path: 'u',
        loadChildren: () => import('../users/feature/user-shell/user-shell.module').then(m => m.UserShellModule)
    },
    {path: '**', redirectTo: 'profile'},
];

@NgModule({
    declarations: [],
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class CoreRoutingModule {
}
