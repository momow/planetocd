Rails.application.routes.draw do

  # scope ":locale", locale: /fr/ do
    root to: "application#index"
    resources :articles
    resources :likes
    get 'about', to: 'static_pages#about', as: :about 
  # end
  
  get '*other_path', to: 'base#not_found'

end
