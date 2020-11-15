class ArticlesController < BaseController
    def index
        redirect_to(root_url)
        #render 'articles/index'
    end

    def show
        param = params[:id]
        param or return not_found

        split = param.split('-')
        id = split[0].to_i
        article = Rails.application.articles[I18n.locale.to_s][id]
        
        article or return not_found
        
        if param != article.to_param
            redirect_to article_url(article)
        end

        params[:article] = article
    end
end
