import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TerminalComponent } from './terminal/terminal.component';
import { NavbarComponent } from './navbar/navbar.component';
import { NavbarButtonComponent } from './navbar-button/navbar-button.component';
import { CodeExampleComponent } from './code-example/code-example.component';
import { LaptopComponent } from './laptop/laptop.component';

@NgModule({
  declarations: [
    AppComponent,
    TerminalComponent,
    NavbarComponent,
    NavbarButtonComponent,
    CodeExampleComponent,
    LaptopComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
