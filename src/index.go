package components

import (
  about "bradswhite.com/src/about"
  work "bradswhite.com/src/work"
  blog "bradswhite.com/src/blog"
  post "bradswhite.com/src/post"
  contact "bradswhite.com/src/contact"
  error "bradswhite.com/src/error"
  home "bradswhite.com/src/home"
  page "bradswhite.com/src/page"
)

/*AboutPage = about.AboutPage
WorkPage = work.WorkPage
BlogPage = blog.BlogPage
PostPage = post.PostPage
ContactPage = contact.ContactPage
ErrorPage = error.ErrorPage
HomePage = home.HomePage
Page = page.Page*/

var AboutPage, WorkPage, BlogPage, PostPage, ContactPage, ErrorPage, HomePage, Page = about.AboutPage, work.WorkPage, blog.BlogPage, post.PostPage, contact.ContactPage, error.ErrorPage, home.HomePage, page.Page
