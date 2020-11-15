class MakeLikesLanguageNotNullable < ActiveRecord::Migration[5.2]
  def change
    Like.where(language: nil).update_all(language: "fr")
    change_column :likes, :language, :string, :null => false 
  end
end
